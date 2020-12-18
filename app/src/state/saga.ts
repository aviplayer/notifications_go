import {
  all,
  call,
  spawn,
  AllEffect,
  ForkEffect,
} from '@redux-saga/core/effects';
import { select } from 'redux-saga/effects';
import { app } from './app/app.saga';

function* rootSaga(): Generator<AllEffect<ForkEffect<void>>, void, void> {
  const sagas = [app,];

  yield all(
    sagas.map((saga) =>
      spawn(function* safeSaga() {
        while (true) {
          try {
            yield call(saga);
            break;
          } catch (e) {
            console.error(e);
          }
        }
      })
    )
  );
}

export { rootSaga };

