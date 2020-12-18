import {createReducer, PayloadAction} from '@reduxjs/toolkit';
import {AppState, ContentType} from './app.types';
import {appStarted} from './app.actions';

export const initialState = {
  path: '/',
  contentType: ContentType.NotStarted,
  messages: []
};



export const app = createReducer<AppState>(initialState, {
  [appStarted.type]: (state) => ({
    ...state,
    contentType: ContentType.Empty,
    path: "/"
  })
});
