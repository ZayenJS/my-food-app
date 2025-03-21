import { useSelector } from 'react-redux';
import { State, useAppDispatch } from '../store';
import { useEffect } from 'react';
import { getBrandsList } from '../store/actions/brand.action';

export const useBrands = () => {
    const { list } = useSelector((state: State) => state.brands);
    const appDispatch = useAppDispatch();

    useEffect(() => {
        if (!list.length) {
            appDispatch(getBrandsList());
        }
    }, [list]);

    return [list];
};
