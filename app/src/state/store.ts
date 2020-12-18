import createSagaMiddleware from 'redux-saga';
import {configureStore, getDefaultMiddleware, Store} from '@reduxjs/toolkit';
import {reducer} from './reducer';
import {rootSaga} from './saga';
import {initialState} from './app/app.reducer';

const sagaMiddleware = createSagaMiddleware();
const middleware = [...getDefaultMiddleware({thunk: false}), sagaMiddleware];

export const createStore = (): Store => {
  const store = configureStore({
    preloadedState: {
      app: {
        ...initialState,
      },
    },
    reducer,
    middleware,
    devTools: process.env.NODE_ENV !== 'production',
  });

  sagaMiddleware.run(rootSaga);
  return store;
};
