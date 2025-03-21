import { FC, useEffect } from 'react';
import { PageLoader } from '../../components/PageLoader/PageLoader';
import { usePageLoading } from '../../hooks/usePageLoading';

export const ShoppingLists: FC = () => {
    const [pageLoading, setPageLoading] = usePageLoading();

    useEffect(() => {
        // TODO: fetch shopping list
        setPageLoading(false);
    }, []);

    return pageLoading ? <PageLoader /> : <div>Listes de courses</div>;
};
