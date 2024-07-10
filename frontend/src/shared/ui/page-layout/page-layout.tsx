import { Box, Flex } from '@chakra-ui/react';
import { FC, ReactElement } from 'react';

interface IProps {
  header?: ReactElement;
  content: ReactElement;
}

export const PageLayout: FC<IProps> = ({ header, content }) => {
  return (
    <Flex flexDirection="column" px={4} py={4}>
      {header}
      <Box flex={1} as="main" pt={8}>
        {content}
      </Box>
    </Flex>
  );
};
