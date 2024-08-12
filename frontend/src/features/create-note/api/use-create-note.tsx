import { useMutation } from '@tanstack/react-query';
import { AxiosResponse } from 'axios';
import { CreateNoteParams, Note, notesController } from 'entities/note';
import { IApiError } from 'shared/api';

export const useCreateNote = () => {
  return useMutation<AxiosResponse<Note>, IApiError, CreateNoteParams>({
    mutationFn: notesController.createNote,
  });
};
