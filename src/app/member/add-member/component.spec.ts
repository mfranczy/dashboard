// Copyright 2020 The Kubermatic Kubernetes Platform contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {ComponentFixture, fakeAsync, flush, TestBed, tick, waitForAsync} from '@angular/core/testing';
import {MatDialogRef} from '@angular/material/dialog';
import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {fakeProject} from '@app/testing/fake-data/project';
import {MatDialogRefMock} from '@app/testing/services/mat-dialog-ref-mock';
import {ProjectMockService} from '@app/testing/services/project-mock';
import {NotificationService} from '@core/services/notification';
import {ProjectService} from '@core/services/project';
import {SharedModule} from '@shared/module';
import {Group} from '@shared/utils/member-utils/member-utils';
import {AddMemberComponent} from './component';
import {MemberService} from '@core/services/member';
import {MemberServiceMock} from '@app/testing/services/member-mock';

const modules: any[] = [BrowserModule, BrowserAnimationsModule, SharedModule];

describe('AddProjectComponent', () => {
  let fixture: ComponentFixture<AddMemberComponent>;
  let component: AddMemberComponent;

  beforeEach(
    waitForAsync(() => {
      TestBed.configureTestingModule({
        imports: [...modules],
        declarations: [AddMemberComponent],
        providers: [
          {provide: MatDialogRef, useClass: MatDialogRefMock},
          {provide: ProjectService, useClass: ProjectMockService},
          {provide: MemberService, useClass: MemberServiceMock},
          NotificationService,
        ],
        teardown: {destroyAfterEach: false},
      }).compileComponents();
    })
  );

  beforeEach(
    waitForAsync(() => {
      fixture = TestBed.createComponent(AddMemberComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
    })
  );

  it(
    'should create the add member component',
    waitForAsync(() => {
      expect(component).toBeTruthy();
    })
  );

  it('form invalid after creating', () => {
    expect(component.addMemberForm.valid).toBeFalsy();
  });

  it('should have required fields', () => {
    expect(component.addMemberForm.valid).toBeFalsy();
    expect(component.addMemberForm.controls.email.valid).toBeFalsy();
    expect(component.addMemberForm.controls.email.hasError('required')).toBeTruthy();
    expect(component.addMemberForm.controls.group.valid).toBeFalsy();
    expect(component.addMemberForm.controls.group.hasError('required')).toBeTruthy();

    component.addMemberForm.controls.email.patchValue('john@doe.com');
    expect(component.addMemberForm.controls.email.hasError('required')).toBeFalsy();
    component.addMemberForm.controls.group.patchValue(Group.Editor);
    expect(component.addMemberForm.controls.group.hasError('required')).toBeFalsy();
  });

  it('should call addMember method', fakeAsync(() => {
    const spy = jest.spyOn(fixture.debugElement.injector.get(MemberService) as any, 'add');

    component.project = fakeProject();
    component.addMemberForm.controls.email.patchValue('john@doe.com');
    component.addMemberForm.controls.group.patchValue('editors');
    component.addMember();
    tick();
    flush();

    expect(spy).toHaveBeenCalled();
  }));
});
