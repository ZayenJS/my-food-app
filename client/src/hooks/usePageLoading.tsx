import { useSelector } from 'react-redux';
import { State, useAppDispatch } from '../store';
import { setPageLoading } from '../store/actions/global.action';

export const usePageLoading = () => {
    const pageLoading = useSelector((state: State) => state.global.pageLoading);
    const appDispatch = useAppDispatch();

    return [
        pageLoading,
        (loading: boolean) => {
            appDispatch(setPageLoading(loading));
        },
    ] as const;
};
