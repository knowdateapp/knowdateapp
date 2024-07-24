import {
  chakra,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Button,
  Modal,
  ModalOverlay,
  ModalHeader,
  ModalContent,
  ModalBody,
  ModalFooter,
  useDisclosure,
} from '@chakra-ui/react';
import { FC } from 'react';
import { useTranslation } from 'react-i18next';
import { useForm } from 'react-hook-form';
import { generatePath, useNavigate } from 'react-router-dom';
import { useAuth } from 'entities/session';
import { Routes } from 'shared/config';
import { useCreateCard } from 'features/create-card/api';

type CreateCardFormData = {
  question: string;
  answer: string;
};

export const CreateCardFormModal: FC = () => {
  const navigate = useNavigate();
  const { t } = useTranslation();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { register, handleSubmit, reset } = useForm<CreateCardFormData>();
  const { mutate } = useCreateCard();
  const { workspace } = useAuth();

  const onClick = () => {
    reset();
    onOpen();
  };

  const onSubmit = (data: CreateCardFormData) => {
    mutate(
      {
        workspace: workspace,
        question: data.question,
        answer: data.answer,
      },
      {
        onError: (error) => console.log('Create card error - ', error),
        onSuccess: () => navigate(generatePath(Routes.Cards)),
      },
    );
    onClose();
  };

  return (
    <>
      <Button onClick={onClick}>{t('features.createCard.createButtonText')}</Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>{t('features.createCard.modalHeader')}</ModalHeader>
          <ModalBody>
            <chakra.form id="create-card-form" onSubmit={handleSubmit(onSubmit)}>
              <Flex flexDirection="column" gap={2}>
                <FormControl>
                  <FormLabel>{t('features.createCard.formInputQuestionLabel')}</FormLabel>
                  <Input
                    {...register('question')}
                    required={true}
                    placeholder={t('features.createCard.formInputQuestionPlaceholder')}
                  />
                </FormControl>
                <FormControl>
                  <FormLabel>{t('features.createCard.formInputAnswerLabel')}</FormLabel>
                  <Input
                    {...register('answer')}
                    required={true}
                    placeholder={t('features.createCard.formInputAnswerPlaceholder')}
                  />
                </FormControl>
              </Flex>
            </chakra.form>
          </ModalBody>
          <ModalFooter>
            <Flex gap={2}>
              <Button type="submit" form="create-card-form">
                {t('features.createCard.saveButtonText')}
              </Button>
              <Button onClick={onClose} colorScheme="gray">
                {t('features.createCard.closeButtonText')}
              </Button>
            </Flex>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};
