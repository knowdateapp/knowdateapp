import { Typography } from 'antd';
import { FC } from 'react';
import { Header, PageLayout } from 'shared/ui';

export const MainPage: FC = () => {
  return (
    <PageLayout header={<Header />}>
      <Typography>main page</Typography>
    </PageLayout>
  );
};
