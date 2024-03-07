import { QueryClientProvider } from '@tanstack/react-query';
import { ConfigProvider } from 'antd';
import { FC } from 'react';
import { RouterProvider } from 'react-router-dom';
import { queryClient } from 'shared/api/query-client.ts';
import { router } from './config/router.tsx';
import { EmotionThemeProvider } from './providers/emotion-theme-provider.tsx';

export const App: FC = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <ConfigProvider theme={{ cssVar: true }}>
        <EmotionThemeProvider>
          <RouterProvider router={router} />
        </EmotionThemeProvider>
      </ConfigProvider>
    </QueryClientProvider>
  );
};
