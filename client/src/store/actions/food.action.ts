import { createAction, createAsyncThunk } from '@reduxjs/toolkit';
import { apiRouter, ApiRoutes } from '../../router/apiRouter';
import { FoodModel } from '../../models/food';
import { Letter } from '../../@types';

enum FoodActions {
  GET_LIST = 'food/getList',
  CREATE = 'food/create',
  DELETE = 'food/delete',
}

export const getFoodsList = createAsyncThunk(FoodActions.GET_LIST, async () => {
  const url = apiRouter.buildLink(ApiRoutes.foods, {
    query: { 'per-page': '500', 'index-by': 'letter' },
  });

  const response = await fetch(url);
  return response.json();
});

export const createFood = createAction(FoodActions.CREATE, (indexedFood: Record<Letter, FoodModel>) => ({
  payload: indexedFood,
}));

export const deleteFood = createAction(FoodActions.DELETE, (foodId: number) => ({
  payload: foodId,
}));
