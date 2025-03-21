import { createReducer } from '@reduxjs/toolkit';
import { createFood, getFoodsList } from '../actions/food.action';
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
            console.log({
                letter,
                state,
                payload: action.payload,
                stateLetter: state.list[letter as Letter],
                payloadLetter: action.payload[letter as Letter],
            });

            state.list[letter as Letter]?.push(action.payload[letter as Letter]);

            state.list[letter as Letter] = state.list[letter as Letter]?.sort((a, b) => a.name.localeCompare(b.name));
        }),
);
