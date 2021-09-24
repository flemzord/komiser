import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { TrendModule } from 'ngx-trend';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { ComputeComponent } from './compute/compute.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { StorageComponent } from './storage/storage.component';
import { NetworkComponent } from './network/network.component';
import { SecurityComponent } from './security/security.component';
import { DataAndAiComponent } from './data-and-ai/data-and-ai.component';

import { AwsService } from './aws.service';
import { StoreService } from './store.service';
import { PaginationModule } from 'ngx-bootstrap/pagination';
import { ProfileComponent } from './profile/profile.component';
import { LimitsComponent } from './limits/limits.component';
import { AwsDashboardComponent } from './dashboard/aws/aws.component';
import { AwsComputeComponent } from './compute/aws/aws.component';
import { AwsStorageComponent } from './storage/aws/aws.component';
import { AwsNetworkComponent } from './network/aws/aws.component';
import { AwsSecurityComponent } from './security/aws/aws.component';
import { AwsDataAndAIComponent } from './data-and-ai/aws/aws.component';
import { AwsLimitsComponent } from './limits/aws/aws.component';
import { AwsProfileComponent } from './profile/aws/aws.component';
import { NotificationsComponent } from './notifications/notifications.component';

const appRoutes: Routes = [
  {
    path: 'compute',
    component: ComputeComponent,
    data: { title: 'Compute - Komiser' }
  },
  {
    path: 'storage',
    component: StorageComponent,
    data: { title: 'Storage - Komiser' }
  },
  {
    path: 'network',
    component: NetworkComponent,
    data: { title: 'Network - Komiser' }
  },
  {
    path: 'security',
    component: SecurityComponent,
    data: { title: 'Security - Komiser' }
  },
  {
    path: 'data-and-ai',
    component: DataAndAiComponent,
    data: { title: 'Data & AI - Komiser' }
  },
  {
    path: 'profile',
    component: ProfileComponent,
    data: { title: 'Profile - Komiser' }
  },
  {
    path: 'limits',
    component: LimitsComponent,
    data: { title: 'Service Limits Checks - Komiser' }
  },
  {
    path: 'notifications',
    component: NotificationsComponent,
    data: { title: 'Notifications - Komiser' }
  },
  { path: '',
    component: DashboardComponent,
    data: { title: 'Dashboard - Komiser' }
  }
];

@NgModule({
  declarations: [
    AppComponent,
    ComputeComponent,
    DashboardComponent,
    StorageComponent,
    NetworkComponent,
    SecurityComponent,
    DataAndAiComponent,
    ProfileComponent,
    LimitsComponent,
    AwsDashboardComponent,
    AwsComputeComponent,
    AwsStorageComponent,
    AwsNetworkComponent,
    AwsSecurityComponent,
    AwsDataAndAIComponent,
    AwsLimitsComponent,
    AwsProfileComponent,
    NotificationsComponent,
  ],
  imports: [
    RouterModule.forRoot(
      appRoutes
    ),
    HttpModule,
    BrowserModule,
    PaginationModule.forRoot(),
    BrowserAnimationsModule,
    TrendModule,
    FormsModule
  ],
  providers: [
    AwsService,
    StoreService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(){}
}
