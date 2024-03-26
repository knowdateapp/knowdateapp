import styled from '@emotion/styled';
import { Layout } from 'antd';
import { FC, PropsWithChildren, ReactElement } from 'react';

const StyledLayout = styled(Layout)`
  height: 100vh;

  main {
    height: 100%;
    padding: ${({ theme }) => theme.padding}px ${({ theme }) => theme.paddingXL}px;
  }
`;

interface IProps extends PropsWithChildren {
  header?: ReactElement;
}

export const PageLayout: FC<IProps> = ({ children, header }) => {
  return (
    <StyledLayout>
      {header}
      <Layout.Content>{children}</Layout.Content>
    </StyledLayout>
  );
};
