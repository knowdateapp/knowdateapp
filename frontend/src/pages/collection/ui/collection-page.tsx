import { Typography } from 'antd';
import { FC } from 'react';
import { Header, PageLayout } from 'shared/ui';

export const CollectionPage: FC = () => {
  return (
    <PageLayout header={<Header />}>
      <Typography>collection page</Typography>
    </PageLayout>
  );
};
