import { Typography } from 'antd';
import { FC } from 'react';
import { Header, PageLayout } from 'shared/ui';

export const TopicPage: FC = () => {
  return (
    <PageLayout header={<Header />}>
      <Typography>topic page</Typography>
    </PageLayout>
  );
};
