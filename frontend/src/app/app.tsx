import { Global } from '@emotion/react';
import { QueryClientProvider } from '@tanstack/react-query';
import { ConfigProvider } from 'antd';
import { FC } from 'react';
import { RouterProvider } from 'react-router-dom';
import { queryClient } from 'shared/api';
import { theme } from 'shared/config';
import { EmotionThemeProvider } from './providers';
import { router } from './router';
import { styles } from './styles.ts';

export const App: FC = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <ConfigProvider theme={theme}>
        <EmotionThemeProvider>
          <>
            <RouterProvider router={router} />
            <Global styles={styles} />
          </>
        </EmotionThemeProvider>
      </ConfigProvider>
    </QueryClientProvider>
  );
};
