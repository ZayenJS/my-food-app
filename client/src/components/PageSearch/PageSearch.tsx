import { usePageSearch } from '../../hooks/usePageSearch';
import { buildClassName } from '../../util';
import { SearchInput } from '../Form/SearchInput/SearchInput';

import classes from './PageSearch.module.scss';

export const PageSearch = () => {
    const [search, setSearch] = usePageSearch();

    return (
        <div className={classes.page_search}>
            <div className={buildClassName(classes.search_container, 'padded-wrapper-1rem')}>
                <SearchInput placeholder="Rechercher" value={search} onChange={setSearch} />
            </div>
            <hr />
        </div>
    );
};
