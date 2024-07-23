import { extendBaseTheme, withDefaultColorScheme } from '@chakra-ui/react';
import { theme as chakraTheme } from '@chakra-ui/theme';

const { Button, Heading, Input, Form, FormError, FormLabel, Card, Modal } = chakraTheme.components;

export const theme = extendBaseTheme(
  {
    styles: {
      global: {
        body: {
          bg: '#f9fafb',
        },
      },
    },
    components: {
      Button,
      Heading,
      Input,
      Form,
      FormError,
      FormLabel,
      Card,
      Modal,
    },
  },
  withDefaultColorScheme({ colorScheme: 'blue' }),
) as Record<string, unknown>;
