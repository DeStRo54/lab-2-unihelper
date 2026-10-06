type ApiRequestConfig = import('axios').AxiosRequestConfig;

type AxiosRequestConfig<Params = undefined> = Params extends undefined
  ? { config?: ApiRequestConfig }
  : { params: Params; config?: ApiRequestConfig };
