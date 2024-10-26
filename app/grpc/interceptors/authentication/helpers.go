package authentication

import (
	"context"
	"crspy2/licenses/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

var UnauthedRoutes = []string{
	"/protofiles.Auth/Login",
}

func authorizeSession(ctx context.Context) (context.Context, error) {
	session, ctx, err := RetrieveSessionFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Authentication failed: %v", err)
	}

	if !session.Staff.Approved {
		return nil, status.Errorf(codes.PermissionDenied, "Your account has not yet been approved")
	}
	return ctx, nil
}

func RetrieveSessionFromContext(ctx context.Context) (*database.SessionModal, context.Context, error) {
	// Extract metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, status.Errorf(codes.InvalidArgument, "Missing metadata")
	}

	// Get session token from metadata
	sessionTokens, ok := md["session_token"]
	if !ok {
		return nil, nil, status.Errorf(codes.Unauthenticated, "Missing session token")
	}

	// Retrieve the client IP from the peer info
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, nil, status.Errorf(codes.Unauthenticated, "Unable to determine client IP")
	}

	clientIP, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "Invalid client IP format")
	}

	// Fetch the session from the database
	session, err := database.Client.Session.Get(sessionTokens[0], clientIP)
	if err != nil {
		return nil, nil, status.Errorf(codes.Unauthenticated, "Invalid session token")
	}

	// Check if the session is expired
	if session.ExpiresAt.Before(time.Now()) {
		return nil, nil, status.Errorf(codes.Unauthenticated, "Session expired")
	}

	ctx = context.WithValue(ctx, "session", session)

	return session, ctx, nil
}

func checkCSRFToken(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Missing metadata")
	}

	csrfToken, ok := md["csrf_token"]
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Missing CSRF token")
	}

	isCsrfTokenInvalid := validateCSRFToken(ctx, csrfToken[0]) != nil
	if isCsrfTokenInvalid {
		return status.Errorf(codes.PermissionDenied, "CSRF token validation failed")
	}
	return nil
}

func validateCSRFToken(ctx context.Context, csrfToken string) error {
	session, _, err := RetrieveSessionFromContext(ctx)
	if err != nil {
		return err
	}

	if session.CsrfToken != csrfToken {
		return status.Errorf(codes.PermissionDenied, "Invalid CSRF token")
	}

	return nil
}
