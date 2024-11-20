import { FC, useEffect, useState } from 'react';
import { ApiRoutes, apiRouter } from '../../router/apiRouter';
import { FoodModelIndexedByLetter } from '../../models/food';
import { FoodsIcon } from '../../components/Icon/FoodsIcon/FoodsIcon';

import classes from './Foods.module.scss';
import { Link } from 'react-router-dom';
import { MoreVerticalIcon } from '../../components/Icon/moreVerticalIcon/MoreVerticalIcon';
import { buildClassName } from '../../util';
import { PlusIcon } from '../../components/Icon/PlusIcon/PlusIcon';
import { clientRouter } from '../../router/clientRouter';
import { PageSearch } from '../../components/PageSearch/PageSearch';

export const Foods: FC = () => {
    const [foods, setFoods] = useState<FoodModelIndexedByLetter>({});

    useEffect(() => {
        if (Object.keys(foods).length > 0) return;

        const url = apiRouter.buildLink(ApiRoutes.foods, {
            query: { limit: '20', 'index-by': 'letter' },
        });

        fetch(url)
            .then((response) => response.json())
            .then((data: FoodModelIndexedByLetter) => setFoods(data));
    }, [foods]);

    return (
        <div>
            <PageSearch />
            <div className={buildClassName('padded-wrapper-1rem', classes.foods_title_container)}>
                <h2 className="page_title">
                    <FoodsIcon className="icon-lg" />
                    <span>Aliments</span>
                </h2>
                <Link className={buildClassName('button')} to={clientRouter.routes.food_new.path}>
                    <PlusIcon className={buildClassName('icon-xs', classes.foods_new_food_icon)} />
                    <span>Créer</span>
                </Link>
            </div>
            <ul className={classes.foods_list}>
                {Object.entries(foods).map(([letter, foods]) => (
                    <li className={classes.foods_list_letter} key={letter}>
                        <h3>{letter}</h3>
                        <hr />
                        <ul className={classes.foods_list_food}>
                            {foods.map((food) => (
                                <li key={food.id}>
                                    <Link to="">{food.name}</Link>
                                    <button tabIndex={0} className={buildClassName(classes.feedback, 'feedback')}>
                                        <MoreVerticalIcon className="icon-sm" />
                                    </button>
                                </li>
                            ))}
                        </ul>
                    </li>
                ))}
            </ul>
        </div>
    );
};
