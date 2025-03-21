import { createAsyncThunk } from '@reduxjs/toolkit';
import { apiRouter, ApiRoutes } from '../../router/apiRouter';

enum BrandActions {
    GET_LIST = 'brand/getList',
}

export const getBrandsList = createAsyncThunk(BrandActions.GET_LIST, async () => {
    const response = await fetch(apiRouter.buildLink(ApiRoutes.brands));
    return response.json();
});
