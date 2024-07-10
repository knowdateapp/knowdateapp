import {
  Heading,
  Stack,
  FormControl,
  FormLabel,
  Input,
  Button,
  Box,
  FormErrorMessage,
  chakra,
} from '@chakra-ui/react';
import { zodResolver } from '@hookform/resolvers/zod';
import { FC } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useTranslation } from 'react-i18next';
import { z } from 'zod';
import { useAuth } from 'entities/session';
import { TranslationKey } from 'shared/config';
import { isAlphaNumeric } from 'shared/lib';
import {
  MAX_PASSWORD_LENGTH,
  MAX_USERNAME_LENGTH,
  MIN_PASSWORD_LENGTH,
  MIN_USERNAME_LENGTH,
} from '../../model';

const schema = z.object({
  username: z
    .string({ required_error: 'errors.required' })
    .min(1, 'errors.required')
    .min(3, 'errors.minMaxlength')
    .max(30, 'errors.minMaxlength')
    .refine(isAlphaNumeric, {
      message: 'errors.onlyAlphaNumeric',
    }),
  password: z
    .string({ required_error: 'errors.required' })
    .min(1, 'errors.required')
    .min(8, 'errors.minMaxlength')
    .max(30, 'errors.minMaxlength'),
});

type FormFields = z.infer<typeof schema>;

export const AuthForm: FC = () => {
  const { t } = useTranslation('translation');
  const { setWorkspace } = useAuth();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormFields>({ resolver: zodResolver(schema) });

  const onSubmit: SubmitHandler<FormFields> = ({ username }) => setWorkspace(username);

  return (
    // TODO: Что за дичь? void handleSubmit(onSubmit)(event)}
    <chakra.form minW="xl" onSubmit={(event) => void handleSubmit(onSubmit)(event)}>
      <Heading>{t('features.auth.authForm.title')}</Heading>
      <Stack mt={10} borderColor="gray.200" position="relative">
        <Stack spacing={10} zIndex={4} backgroundColor="white" p={12} mx={4} borderRadius="2xl">
          <FormControl isInvalid={Boolean(errors.username)}>
            <FormLabel aria-required="false">{t('username')}</FormLabel>
            <Input {...register('username')} />
            <FormErrorMessage>
              {t(errors.username?.message as TranslationKey, {
                min: MIN_USERNAME_LENGTH,
                max: MAX_USERNAME_LENGTH,
              })}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.password)}>
            <FormLabel>{t('password')}</FormLabel>
            <Input {...register('password')} type="password" />
            <FormErrorMessage>
              {t(errors.password?.message as TranslationKey, {
                min: MIN_PASSWORD_LENGTH,
                max: MAX_PASSWORD_LENGTH,
              })}
            </FormErrorMessage>
          </FormControl>
          <Button type="submit">{t('logIn')}</Button>
        </Stack>
        <Box
          position="absolute"
          inset="-5px"
          transform="rotate(-2deg)"
          backgroundImage="linear-gradient(to top right, cyan.400, blue.500);"
          borderRadius="2xl"
          zIndex={3}
        />
        <Box
          position="absolute"
          inset="-6px"
          zIndex={2}
          borderRadius="lg"
          backgroundColor="gray.100"
        />
      </Stack>
    </chakra.form>
  );
};
