import { RootState } from '../types';
import { ContentType } from './app.types';


export const getPath = (state: RootState): string => state.app.path;

export const getContentType = (state: RootState): ContentType =>
  state.app.contentType;

