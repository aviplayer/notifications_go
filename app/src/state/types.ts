import {AppState} from './app/app.types';

export interface RootState {
  app: AppState;
}

export interface NavigationMeta {
  isStartup: boolean;
}

export interface NotificationType {
  id: number
  type: number
  title: string
  description: string
  template: string
  email: string
  pwd: string
  smtp_server: string
  smpt_port: number|string
}

