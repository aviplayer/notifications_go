import React, {useEffect} from 'react';
import {useDispatch, useSelector} from "react-redux";
import {getContentType} from "./state/app/app.selectors";
import {appStart} from "./state/app/app.actions";
import {AppContainer} from "./containers/AppContainer";

export const App: React.FC = React.memo(() => {
  const dispatch = useDispatch();
  const contentType = useSelector(getContentType);
  useEffect(() => {
    dispatch(appStart());
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <AppContainer/>
  );
});

export default App;
