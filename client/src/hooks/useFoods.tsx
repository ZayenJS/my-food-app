import { useSelector } from 'react-redux';
import { State, useAppDispatch } from '../store';
import { useEffect } from 'react';
import { getFoodsList } from '../store/actions/food.action';

export const useFoods = () => {
    const { list } = useSelector((state: State) => state.foods);
    const appDispatch = useAppDispatch();

    useEffect(() => {
        if (!Object.keys(list).length) {
            appDispatch(getFoodsList());
        }
    }, [list]);

    return [list];
};
