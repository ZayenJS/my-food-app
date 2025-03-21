import { createReducer } from '@reduxjs/toolkit';
import { getBrandsList } from '../actions/brand.action';
import { Brand } from '../../models/brand';

export type BrandState = { list: Brand[] };

const INITIAL_STATE: BrandState = {
    list: [],
};

export const brandReducer = createReducer(INITIAL_STATE, (builder) =>
    builder.addCase(getBrandsList.fulfilled, (state, action) => {
        state.list = action.payload;
    }),
);
