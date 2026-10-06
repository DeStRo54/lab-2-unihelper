import { api } from '@/utils/api/instance';

export type PostAdminAddGroupParams = {
  name: string;
  course: number;
  icalLink: string;
};

export type PostAdminAddGroupConfig = AxiosRequestConfig<PostAdminAddGroupParams>;

export const postAdminAddGroup = async ({ params, config }: PostAdminAddGroupConfig) =>
  api.post<AdminAddGroupResponse>('/admin/addGroup', params, config);
