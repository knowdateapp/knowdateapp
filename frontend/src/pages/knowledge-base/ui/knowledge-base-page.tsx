import { Typography } from 'antd';
import { FC } from 'react';
import { Header, PageLayout } from 'shared/ui';

export const KnowledgeBasePage: FC = () => {
  return (
    <PageLayout header={<Header />}>
      <Typography>knowledgebase page</Typography>
    </PageLayout>
  );
};
