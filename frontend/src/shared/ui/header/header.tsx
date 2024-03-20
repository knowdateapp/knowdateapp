import styled from '@emotion/styled';
import { Layout, Typography } from 'antd';
import { FC } from 'react';

const StyledText = styled(Typography.Title)`
  && {
    color: ${({ theme }) => theme.colorWhite};
    margin: 0;
  }
`;

const StyledHeader = styled(Layout.Header)`
  display: flex;
  align-items: center;
`;

export const Header: FC = () => {
  return (
    <StyledHeader>
      <StyledText level={1}>KnowlegdeHub</StyledText>
    </StyledHeader>
  );
};
