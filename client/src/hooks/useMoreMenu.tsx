import { useSelector } from 'react-redux';
import { State, useAppDispatch } from '../store';
import { setOpenedMoreMenu } from '../store/actions/global.action';

export const useMoreMenu = () => {
  const openedMoreMenu = useSelector((state: State) => state.global.openedMoreMenu);
  const appDispatch = useAppDispatch();

  return [
    openedMoreMenu,
    (id: string | null) => {
      appDispatch(setOpenedMoreMenu(id));
    },
  ] as const;
};
