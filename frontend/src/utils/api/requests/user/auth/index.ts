import { api } from '@/utils/api/instance';

export type PostUserAuthParams = {
  email: string;
  password: string;
};

export type PostUserAuthConfig = AxiosRequestConfig<PostUserAuthParams>;

export const postUserAuth = async ({ params, config }: PostUserAuthConfig) =>
  api.post<UserAuthResponse>('/user/auth', params, config);
