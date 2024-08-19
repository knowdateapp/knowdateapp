import { ArrowLeftIcon, ArrowRightIcon } from '@chakra-ui/icons';
import { Box, Text } from '@chakra-ui/react';

type EditorHeaderProps = {
  isSidebarOpen?: boolean;
  toggleSidebar?: () => void;
  characters: number;
  words: number;
};

export const EditorHeader = ({
  characters,
  words,
  isSidebarOpen,
  toggleSidebar,
}: EditorHeaderProps) => {
  return (
    <Box
      display="flex"
      justifyContent="space-between"
      flex="none"
      color="black"
      backgroundColor="white"
      py={2}
      pl={6}
      pr={3}
      border="1px solid black"
    >
      <Box display="flex" alignItems="center" gap={2}>
        <Box display="flex" alignItems="center" gap={2}>
          {isSidebarOpen ? (
            <ArrowLeftIcon onClick={toggleSidebar} />
          ) : (
            <ArrowRightIcon onClick={toggleSidebar} />
          )}
        </Box>
      </Box>
      <Box display="flex" alignItems="center">
        <Box
          display="flex"
          flexDirection="column"
          pr={4}
          mr={4}
          textAlign="right"
          borderRight="1px solid black"
        >
          <Text fontSize="sm">
            {words} {words === 1 ? 'word' : 'words'}
          </Text>
          <Text fontSize="sm">
            {characters} {characters === 1 ? 'character' : 'characters'}
          </Text>
        </Box>
      </Box>
    </Box>
  );
};
