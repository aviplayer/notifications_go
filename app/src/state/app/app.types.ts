import {Action} from "redux";

export interface AppState {
  path: string;
  contentType: ContentType;
  messages: string[];
}

export enum ContentType {
  NotStarted,
  Empty
}

export interface NotificationSelector {
  id: number;
  title: string;
  description: string;
}

