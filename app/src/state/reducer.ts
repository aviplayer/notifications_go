import { combineReducers, Reducer } from 'redux';
import { enableMapSet } from 'immer';
import { RootState } from './types';
import { app } from './app/app.reducer';

enableMapSet();

export const reducer: Reducer<RootState> = combineReducers({
  app,
});
