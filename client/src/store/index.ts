import { Action, configureStore, ThunkAction } from '@reduxjs/toolkit';
import { useDispatch } from 'react-redux';

import rootReducer from './reducers';

const store = configureStore({
    reducer: rootReducer,
    devTools: true,
});

export type State = ReturnType<typeof rootReducer>;
export type Thunk<ReturnType = void> = ThunkAction<ReturnType, State, unknown, Action>;
export type AppDispatch = typeof store.dispatch;
export const useAppDispatch = () => useDispatch<AppDispatch>();

export default store;
