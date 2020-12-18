import {AxiosResponse} from 'axios';
import {httpClient} from '../../utils';

export const getNotifications = (): Promise<AxiosResponse<AxiosResponse>> =>
  httpClient.get('notifications', {
      headers: {
        Accept: 'application/json',
      },
    }
  );
