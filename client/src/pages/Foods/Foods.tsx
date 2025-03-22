import { FC, useEffect, useState } from 'react';
import { FoodModelIndexedByLetter } from '../../models/food';
import { FoodsIcon } from '../../components/Icon/FoodsIcon/FoodsIcon';

import classes from './Foods.module.scss';
import { Link } from 'react-router-dom';
import { buildClassName } from '../../util';
import { PlusIcon } from '../../components/Icon/PlusIcon/PlusIcon';
import { clientRouter } from '../../router/clientRouter';
import { PageSearch } from '../../components/PageSearch/PageSearch';
import { PageLoader } from '../../components/PageLoader/PageLoader';
import { usePageLoading } from '../../hooks/usePageLoading';
import { usePageSearch } from '../../hooks/usePageSearch';
import { Letter } from '../../@types';
import { useFoods } from '../../hooks/useFoods';
import { MoreMenu } from '../../components/MoreMenu/MoreMenu';

export const Foods: FC = () => {
  const [foods] = useFoods();
  const [filteredFoods, setFilteredFoods] = useState<FoodModelIndexedByLetter>({});
  const [pageLoading, setPageLoading] = usePageLoading();
  const [pageSearch] = usePageSearch();

  useEffect(() => {
    if (pageSearch === '') {
      setFilteredFoods(foods);
      return;
    }

    const filtered = Object.entries(foods).reduce<FoodModelIndexedByLetter>(
      (acc: FoodModelIndexedByLetter, [letter, foods]) => {
        const filteredFoods = foods.filter(food => food.name.toLowerCase().includes(pageSearch.toLowerCase()));

        if (filteredFoods.length > 0) {
          acc[letter as Letter] = filteredFoods;
        }

        return acc;
      },
      {} as FoodModelIndexedByLetter,
    );

    setFilteredFoods(filtered);
  }, [pageSearch, foods]);

  useEffect(() => {
    if (!Object.keys(foods).length && !pageLoading) {
      setPageLoading(true);
    } else if (Object.keys(foods).length && pageLoading) {
      setPageLoading(false);
    }
  }, [foods]);

  return pageLoading ? (
    <PageLoader />
  ) : (
    <>
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
        {Object.entries(filteredFoods).map(([letter, foods]) => (
          <li className={classes.foods_list_letter} key={letter}>
            <h3>{letter.toUpperCase()}</h3>
            <hr />
            <ul className={classes.foods_list_food}>
              {foods.map(food => (
                <li key={food.food_id}>
                  <Link to="" className={classes.food_name}>
                    <span>{food.name}</span>
                  </Link>
                  <MoreMenu classNames={[classes.feedback]} foodId={food.food_id} />
                </li>
              ))}
            </ul>
          </li>
        ))}
      </ul>
    </>
  );
};
