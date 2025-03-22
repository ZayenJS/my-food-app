import { FC, useCallback, useEffect, useId, useState } from 'react';
import { buildClassName } from '../../util';
import { MoreVerticalIcon } from '../Icon/moreVerticalIcon/MoreVerticalIcon';

import classes from './MoreMenu.module.scss';
import { useMoreMenu } from '../../hooks/useMoreMenu';
import { apiRouter, ApiRoutes } from '../../router/router';
import { useAppDispatch } from '../../store';
import { deleteFood } from '../../store/actions/food.action';
import { ConfirmModal } from '../ConfirmModal/ConfirmModal';

interface Props {
  classNames?: string[];
  foodId: number;
}

export const MoreMenu: FC<Props> = ({ classNames, foodId }) => {
  const id = useId();
  const [activeMenuId, setActiveMenuId] = useMoreMenu();

  const [deleteModalVisible, setDeleteModalVisible] = useState(false);

  const appDispatch = useAppDispatch();

  const showMenu = activeMenuId === id;

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

  const onDeleteHandler = async () => {
    try {
      await fetch(apiRouter.buildLink(ApiRoutes.delete_food, { params: { id: foodId.toString() } }), {
        method: 'DELETE',
      });
      appDispatch(deleteFood(foodId));

      return true;
    } catch (e) {
      console.error(e);
      return false;
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
            <button
              type="button"
              onClick={() => {
                setDeleteModalVisible(true);
                onShowMenuHandler();
              }}
            >
              Delete
            </button>
          </li>
        </ul>
      </div>
      <ConfirmModal
        text="Voulez-vous vraiment supprimer cet aliment ?"
        visible={deleteModalVisible}
        onConfirm={onDeleteHandler}
        onCancel={setDeleteModalVisible.bind(this, false)}
      />
    </div>
  );
};
