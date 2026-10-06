import { api } from '@/utils/api/instance';

export type PatchAdminUpdateClasses = AxiosRequestConfig;

export const patchAdminUpdateClasses = async (requestConfig?: PatchAdminUpdateClasses) =>
  api.patch<AdminUpdateClassesResponse>('/admin/updateClasses', requestConfig?.config);
