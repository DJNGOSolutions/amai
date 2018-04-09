import { TestBed, inject } from '@angular/core/testing';

import { AmaiService } from './amai.service';

describe('AmaiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AmaiService]
    });
  });

  it('should be created', inject([AmaiService], (service: AmaiService) => {
    expect(service).toBeTruthy();
  }));
});
