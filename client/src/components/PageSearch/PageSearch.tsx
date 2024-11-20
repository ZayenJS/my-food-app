import { buildClassName } from '../../util';
import { SearchInput } from '../Form/SearchInput/SearchInput';

import classes from './PageSearch.module.scss';

export const PageSearch = () => {
    return (
        <div className={classes.page_search}>
            <div className={buildClassName(classes.search_container, 'padded-wrapper-1rem')}>
                <SearchInput placeholder="Rechercher" />
            </div>
            <hr />
        </div>
    );
};
