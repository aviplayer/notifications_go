import {
  call,
  put,
  select,
  takeEvery,
  takeLatest,
} from '@redux-saga/core/effects';
import { appStart, appStarted } from './app.actions';

import { getPath } from './app.selectors';
import { history } from '../history';
import {getNotifications} from "./app.api";
import {AxiosResponse} from "axios";




function* appStartSaga() {
  try {
    const { data }: AxiosResponse<AxiosResponse> = yield call(getNotifications);
    yield put(appStarted());
  } catch (e) {
    console.error('Application failed: ' + e);
  }
}

function* navigatedByActionSaga() {
  const path = yield select(getPath);
  if (history.location.pathname !== path) {
    history.push(path);
  }
}

export function* app(): Generator {
  yield takeLatest(appStart.type, appStartSaga);
  yield takeLatest(appStarted.type, navigatedByActionSaga);
}
