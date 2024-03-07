import { ThemeProvider } from '@emotion/react';
import { theme } from 'antd';
import { FC, ReactElement } from 'react';

interface IProps {
  children: ReactElement;
}

export const EmotionThemeProvider: FC<IProps> = ({ children }) => {
  const antdTheme = theme.useToken();

  return <ThemeProvider theme={antdTheme.token}>{children}</ThemeProvider>;
};
