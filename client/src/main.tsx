import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import './assets/scss/main.scss';
import { Provider } from 'react-redux';
import store from './store/index.ts';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Layout } from './components/Layout/Layout.tsx';
import { Recipes } from './pages/Recipes/Recipes.tsx';
import { clientRouter, AppRoutes } from './router/router.ts';

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Provider store={store}>
            <BrowserRouter>
                <Layout>
                    <Routes>
                        <Route path={clientRouter.buildLink(AppRoutes.home)} Component={Recipes} />
                    </Routes>
                </Layout>
            </BrowserRouter>
        </Provider>
    </StrictMode>,
);
