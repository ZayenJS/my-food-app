import { createReducer } from '@reduxjs/toolkit';
import { createFood, deleteFood, getFoodsList } from '../actions/food.action';
import { FoodModelIndexedByLetter } from '../../models/food';
import { Letter } from '../../@types';

export type FoodState = { list: FoodModelIndexedByLetter };

const INITIAL_STATE: FoodState = {
  list: {},
};

export const foodReducer = createReducer(INITIAL_STATE, (builder) =>
  builder
    .addCase(getFoodsList.fulfilled, (state, action) => {
      state.list = action.payload;
    })
    .addCase(createFood, (state, action) => {
      const letter = Object.keys(action.payload)[0];

      if (!state.list[letter as Letter]) {
        state.list[letter as Letter] = [];
      }

      state.list[letter as Letter]!.push(action.payload[letter as Letter]);
      state.list[letter as Letter] = state.list[letter as Letter]?.sort((a, b) => a.name.localeCompare(b.name));

      // sort the list of foods by letter
      state.list = Object.keys(state.list)
        .sort((a, b) => a.localeCompare(b))
        .reduce<FoodModelIndexedByLetter>((acc, letter) => {
          acc[letter as Letter] = state.list[letter as Letter];
          return acc;
        }, {});
    })
    .addCase(deleteFood, (state, action) => {
      const letter = Object.keys(state.list).find((letter) =>
        state.list[letter as Letter]?.some((food) => food.food_id === action.payload),
      );

      if (!letter) {
        return;
      }

      state.list[letter as Letter] = state.list[letter as Letter]?.filter((food) => food.food_id !== action.payload);

      if (state.list[letter as Letter]?.length === 0) {
        delete state.list[letter as Letter];
      }
    }),
);
