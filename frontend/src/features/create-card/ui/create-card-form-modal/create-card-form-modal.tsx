import {
  Button,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  chakra,
} from '@chakra-ui/react';
import { FC } from 'react';
import { useForm } from 'react-hook-form';
import { generatePath, useNavigate } from 'react-router-dom';
import { useAuth } from 'entities/session';
import { Routes } from 'shared/config';
import { useCreateCard } from 'features/create-card/api';

type Card = {
  question: string;
  answer: string;
};

type Props = {
  isOpen: boolean;
  onClose: () => void;
};

export const CreateCardFormModal: FC<Props> = ({ isOpen, onClose }) => {
  const { workspace } = useAuth();
  const { register, handleSubmit } = useForm<Card>();
  const { mutate } = useCreateCard();
  const navigate = useNavigate();

  const onSubmit = (card: Card) => {
    mutate(
      {
        workspace: workspace,
        question: card.question,
        answer: card.answer,
      },
      {
        onError: (error) => console.log(`create card error - `, error),
        onSuccess: () => {
          navigate(generatePath(Routes.Cards));
        },
      },
    );
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} scrollBehavior="outside" size="lg">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Создайте карточку</ModalHeader>
        <ModalBody>
          <chakra.form id="create-card-form" onSubmit={handleSubmit(onSubmit)}>
            <Flex flexDirection="column" gap={2}>
              <FormControl>
                <FormLabel>Вопрос</FormLabel>
                <Input {...register('question')} required={true} placeholder="Question" />
              </FormControl>
              <FormControl>
                <FormLabel>Ответ</FormLabel>
                <Input {...register('answer')} required={true} placeholder="Answer" />
              </FormControl>
            </Flex>
          </chakra.form>
        </ModalBody>
        <ModalFooter>
          <Flex gap={2}>
            <Button type="submit" form="create-card-form">
              Сохранить
            </Button>
            <Button onClick={onClose} colorScheme="blue">
              Закрыть
            </Button>
          </Flex>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
