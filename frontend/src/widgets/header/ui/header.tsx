import { Flex, Heading } from '@chakra-ui/react';
import { FC, useEffect, useState } from 'react';
import { AuthButton } from 'features/auth';
import { CreateNoteButton } from 'features/create-note';
import { useAuth } from 'entities/session';

export const Header: FC = () => {
  const { isAuth } = useAuth();
  const [isSticky, setIsSticky] = useState(false);

  useEffect(() => {
    const onScroll = () => {
      setIsSticky(window.scrollY > 20);
    };
    window.addEventListener('scroll', onScroll);

    return () => {
      window.removeEventListener('scroll', onScroll);
    };
  }, []);

  return (
    <Flex
      as="header"
      position="sticky"
      top={4}
      justifyContent="space-between"
      alignItems="center"
      borderRadius="2xl"
      zIndex={10}
      transition="0.3s all ease-in-out"
      p={isSticky ? 4 : 0}
      borderColor="gray.200"
      borderWidth={isSticky ? '1px' : 0}
      bgColor={isSticky ? 'white' : 'transparent'}
    >
      <Heading size="xl">KnowlegdeHub</Heading>
      {isAuth && (
        <Flex gap={2}>
          <CreateNoteButton />
          <AuthButton />
        </Flex>
      )}
    </Flex>
  );
};
