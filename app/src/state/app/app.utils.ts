import { Action } from '@reduxjs/toolkit';


export const parseQueryParams = (
  query: string
): { [param: string]: string } => {
  if (!query || query.length < 2) {
    return {};
  }
  return query
    .substring(1)
    .split('&')
    .map((kv) => kv.split('='))
    .reduce((acc, [key, value]) => ({ ...acc, [key]: value }), {});
};


