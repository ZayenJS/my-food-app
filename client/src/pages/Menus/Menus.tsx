import { FC, useEffect } from 'react';
import { MenuIcon } from '../../components/Icon/MenuIcon/MenuIcon';
import { PageLoader } from '../../components/PageLoader/PageLoader';
import { usePageLoading } from '../../hooks/usePageLoading';

export const Menus: FC = () => {
    const [pageLoading, setPageLoading] = usePageLoading();

    useEffect(() => {
        // TODO: fetch menus
        if (pageLoading) setPageLoading(false);
    }, []);

    return pageLoading ? (
        <PageLoader />
    ) : (
        <div className="padded-wrapper-1rem">
            <h2 className="page_title">
                <MenuIcon className="icon-lg" />
                <span>Menus</span>
            </h2>
        </div>
    );
};
