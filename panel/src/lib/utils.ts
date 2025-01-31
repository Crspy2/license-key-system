import { clsx, type ClassValue } from "clsx"
import { extendTailwindMerge } from 'tailwind-merge'
import { withFluid } from '@fluid-tailwind/tailwind-merge'

export const twMerge = extendTailwindMerge(withFluid)

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))

}

const permissionNames: Record<string, number> = {
  Default: 1 << 0,
  HWIDReset: 1 << 1,
  PasswordReset: 1 << 2,
  Compensate: 1 << 3,
  ProductStatus: 1 << 4,
  ManageProducts: 1 << 5,
  GenerateKeys: 1 << 6,
  ViewLogs: 1 << 7,
  ManageUsers: 1 << 8,
  ManageStaff: 1 << 9,
};

export const ConvertPermissionsToValues = (initialPermissions: string[]) => {
  return initialPermissions
      .map((perm) => permissionNames[perm])
      .filter((value): value is number => value !== undefined);
};


export const GetUserRoleText = (role: number) => {
  switch (role) {
    case 0: return "Staff";
    case 1: return "Senior Staff";
    case 2: return "Lead Staff";
    case 3: return "Dev";
    case 4: return "Owner";
  }
}

export const GetUserRoleValue = (role: string) => {
  switch (role) {
    case "Staff": return 0;
    case "Senior Staff": return 1;
    case "Lead Staff": return 2;
    case "Developer": return 3;
    case "Owner": return 4;
  }
}