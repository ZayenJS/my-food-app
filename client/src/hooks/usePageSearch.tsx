import { useSelector } from 'react-redux';
import { State, useAppDispatch } from '../store';
import { setPageSearch } from '../store/actions/global.action';

export const usePageSearch = () => {
    const search = useSelector((state: State) => state.global.pageSearch);
    const appDispatch = useAppDispatch();

    return [
        search,
        (search: string) => {
            appDispatch(setPageSearch(search));
        },
    ] as const;
};
