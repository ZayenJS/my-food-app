import { FC, useCallback, useEffect, useId, useRef, useState } from 'react';
import { buildClassName } from '../../util';
import { MoreVerticalIcon } from '../Icon/moreVerticalIcon/MoreVerticalIcon';

import classes from './MoreMenu.module.scss';
import { useMoreMenu } from '../../hooks/useMoreMenu';
import { LoadingButton } from '../LoadingButton/LoadingButton';
import { apiRouter, ApiRoutes } from '../../router/router';
import { useAppDispatch } from '../../store';
import { deleteFood } from '../../store/actions/food.action';

interface Props {
  classNames?: string[];
  foodId: number;
}

export const MoreMenu: FC<Props> = ({ classNames, foodId }) => {
  const id = useId();
  const [activeMenuId, setActiveMenuId] = useMoreMenu();
  const appDispatch = useAppDispatch();

  const [loading, setLoading] = useState(false);
  const showMenu = activeMenuId === id;

  const loadingButtonRef = useRef<HTMLButtonElement>(null);

  const clickOutsideHandler = useCallback(
    (e: MouseEvent) => {
      if (e.target instanceof HTMLElement && !e.target.closest(`.${classes.root}`) && showMenu) {
        setActiveMenuId(null);
      }
    },
    [setActiveMenuId, showMenu],
  );

  useEffect(() => {
    document.addEventListener('click', clickOutsideHandler);

    return () => {
      document.removeEventListener('click', clickOutsideHandler);
    };
  }, [clickOutsideHandler]);

  const onShowMenuHandler = () => {
    setActiveMenuId(showMenu ? null : id);
  };

  const changeLoadingButtonStatus = (status: boolean) => {
    const buttonWidth = loadingButtonRef.current?.offsetWidth;
    const buttonHeight = loadingButtonRef.current?.offsetHeight;

    if (!loadingButtonRef.current) {
      setLoading(status);
      return;
    }

    if (status) {
      loadingButtonRef.current.style.width = `${buttonWidth}px`;
      loadingButtonRef.current.style.height = `${buttonHeight}px`;
    } else {
      loadingButtonRef.current.style.width = ``;
      loadingButtonRef.current.style.height = ``;
    }

    setLoading(status);
  };

  const onDeleteHandler = async () => {
    changeLoadingButtonStatus(true);

    try {
      await fetch(apiRouter.buildLink(ApiRoutes.delete_food, { params: { id: foodId.toString() } }), {
        method: 'DELETE',
      });
      appDispatch(deleteFood(foodId));
    } finally {
      changeLoadingButtonStatus(false);
    }
  };

  return (
    <div className={classes.root}>
      <button tabIndex={0} className={buildClassName('feedback', ...(classNames ?? []))} onClick={onShowMenuHandler}>
        <MoreVerticalIcon className="icon-sm" />
      </button>

      <div className={buildClassName(classes.more_menu, showMenu ? classes.show : '')}>
        <ul className={classes.menu}>
          <li>
            <button>Edit</button>
          </li>
          <li>
            <LoadingButton
              type="button"
              ref={loadingButtonRef}
              classNames={{
                loading: buildClassName('loading', classes.loading),
              }}
              loading={loading}
              onClick={onDeleteHandler}>
              Delete
            </LoadingButton>
          </li>
        </ul>
      </div>
    </div>
  );
};
