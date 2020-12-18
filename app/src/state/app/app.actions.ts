import { createAction, PrepareAction } from '@reduxjs/toolkit';
import {NavigationMeta} from "../types";


function prepareNavigationAction<P>(payload: P, isStartup: boolean) {
  const meta: NavigationMeta = {
    isStartup,
  };
  return {
    payload,
    meta,
  };
}

export const appStart = createAction('app/START');
export const appStarted = createAction('app/STARTED');
