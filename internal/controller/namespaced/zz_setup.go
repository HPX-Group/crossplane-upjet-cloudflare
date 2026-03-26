// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accessrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/accessrule"
	zerotrustaccessaicontrolsmcpportal "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessaicontrolsmcpportal"
	zerotrustaccessaicontrolsmcpserver "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessaicontrolsmcpserver"
	zerotrustaccessapplication "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessapplication"
	zerotrustaccesscustompage "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccesscustompage"
	zerotrustaccessgroup "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessgroup"
	zerotrustaccessidentityprovider "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessidentityprovider"
	zerotrustaccessinfrastructuretarget "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessinfrastructuretarget"
	zerotrustaccesskeyconfiguration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccesskeyconfiguration"
	zerotrustaccessmtlscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessmtlscertificate"
	zerotrustaccessmtlshostnamesettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessmtlshostnamesettings"
	zerotrustaccesspolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccesspolicy"
	zerotrustaccessservicetoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessservicetoken"
	zerotrustaccessshortlivedcertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccessshortlivedcertificate"
	zerotrustaccesstag "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustaccesstag"
	zerotrustorganization "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/zerotrustorganization"
	account "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/account"
	accountsubscription "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/accountsubscription"
	accounttoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/accounttoken"
	apitoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/apitoken"
	member "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/member"
	registrardomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/registrardomain"
	apishield "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishield"
	apishielddiscoveryoperation "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishielddiscoveryoperation"
	apishieldoperation "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldoperation"
	apishieldoperationschemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldoperationschemavalidationsettings"
	apishieldschema "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldschema"
	apishieldschemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldschemavalidationsettings"
	schemavalidationoperationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/schemavalidationoperationsettings"
	schemavalidationschemas "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/schemavalidationschemas"
	schemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/schemavalidationsettings"
	argosmartrouting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/argosmartrouting"
	argotieredcaching "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/argotieredcaching"
	magicnetworkmonitoringconfiguration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicnetworkmonitoringconfiguration"
	magicnetworkmonitoringrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicnetworkmonitoringrule"
	magictransitconnector "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magictransitconnector"
	magictransitsite "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magictransitsite"
	magictransitsiteacl "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magictransitsiteacl"
	magictransitsitelan "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magictransitsitelan"
	magictransitsitewan "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magictransitsitewan"
	magicwangretunnel "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicwangretunnel"
	magicwanipsectunnel "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicwanipsectunnel"
	magicwanstaticroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicwanstaticroute"
	regionaltieredcache "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/regionaltieredcache"
	tieredcache "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/tieredcache"
	authenticatedoriginpulls "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/authenticatedoriginpulls"
	authenticatedoriginpullscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/authenticatedoriginpullscertificate"
	authenticatedoriginpullssettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/authenticatedoriginpullssettings"
	certificatepack "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/certificatepack"
	customhostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/customhostname"
	customhostnamefallbackorigin "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/customhostnamefallbackorigin"
	customssl "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/customssl"
	hostnametlssetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/hostnametlssetting"
	keylesscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/keylesscertificate"
	mtlscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/mtlscertificate"
	origincacertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/origincacertificate"
	totaltls "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/totaltls"
	universalsslsetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/universalsslsetting"
	accountdnssettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/accountdnssettings"
	accountdnssettingsinternalview "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/accountdnssettingsinternalview"
	dnsfirewall "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnsfirewall"
	dnsrecord "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnsrecord"
	dnszonetransfersacl "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnszonetransfersacl"
	dnszonetransfersincoming "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnszonetransfersincoming"
	dnszonetransfersoutgoing "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnszonetransfersoutgoing"
	dnszonetransferspeer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnszonetransferspeer"
	dnszonetransferstsig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnszonetransferstsig"
	emailroutingaddress "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingaddress"
	emailroutingcatchall "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingcatchall"
	emailroutingdns "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingdns"
	emailroutingrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingrule"
	emailroutingsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingsettings"
	emailsecurityblocksender "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailsecurityblocksender"
	emailsecurityimpersonationregistry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailsecurityimpersonationregistry"
	emailsecuritytrusteddomains "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailsecuritytrusteddomains"
	botmanagement "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/botmanagement"
	custompages "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/custompages"
	filter "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/filter"
	firewallrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/firewallrule"
	leakedcredentialcheck "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/leakedcredentialcheck"
	leakedcredentialcheckrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/leakedcredentialcheckrule"
	managedtransforms "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/managedtransforms"
	pagerule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/pagerule"
	pageshieldpolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/pageshieldpolicy"
	ratelimit "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/ratelimit"
	ruleset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/ruleset"
	urlnormalizationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/urlnormalizationsettings"
	useragentblockingrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/useragentblockingrule"
	healthcheck "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/healthcheck/healthcheck"
	list "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/list/list"
	listitem "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/list/listitem"
	loadbalancer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/loadbalancer/loadbalancer"
	loadbalancermonitor "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/loadbalancer/loadbalancermonitor"
	loadbalancerpool "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/loadbalancer/loadbalancerpool"
	logpullretention "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/logpush/logpullretention"
	logpushjob "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/logpush/logpushjob"
	logpushownershipchallenge "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/logpush/logpushownershipchallenge"
	addressmap "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/addressmap"
	byoipprefix "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/byoipprefix"
	callssfuapp "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/callssfuapp"
	callsturnapp "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/callsturnapp"
	cloudconnectorrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudconnectorrules"
	cloudforceonerequest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequest"
	cloudforceonerequestasset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequestasset"
	cloudforceonerequestmessage "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequestmessage"
	cloudforceonerequestpriority "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequestpriority"
	connectivitydirectoryservice "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/connectivitydirectoryservice"
	contentscanning "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/contentscanning"
	contentscanningexpression "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/contentscanningexpression"
	d1database "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/d1database"
	hyperdriveconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/hyperdriveconfig"
	image "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/image"
	imagevariant "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/imagevariant"
	observatoryscheduledtest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/observatoryscheduledtest"
	organization "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/organization"
	organizationprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/organizationprofile"
	queue "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/queue"
	queueconsumer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/queueconsumer"
	regionalhostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/regionalhostname"
	snippet "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/snippet"
	snippetrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/snippetrules"
	snippets "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/snippets"
	ssoconnector "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/ssoconnector"
	tokenvalidationconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/tokenvalidationconfig"
	tokenvalidationrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/tokenvalidationrules"
	turnstilewidget "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/turnstilewidget"
	user "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/user"
	web3hostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/web3hostname"
	webanalyticsrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/webanalyticsrule"
	webanalyticssite "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/webanalyticssite"
	notificationpolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/notification/notificationpolicy"
	notificationpolicywebhooks "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/notification/notificationpolicywebhooks"
	providerconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/providerconfig"
	r2bucket "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucket"
	r2bucketcors "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketcors"
	r2bucketeventnotification "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketeventnotification"
	r2bucketlifecycle "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketlifecycle"
	r2bucketlock "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketlock"
	r2bucketsippy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketsippy"
	r2customdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2customdomain"
	r2manageddomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2manageddomain"
	spectrumapplication "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/spectrum/spectrumapplication"
	stream "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/stream"
	streamaudiotrack "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamaudiotrack"
	streamcaptionlanguage "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamcaptionlanguage"
	streamdownload "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamdownload"
	streamkey "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamkey"
	streamliveinput "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamliveinput"
	streamwatermark "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamwatermark"
	streamwebhook "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamwebhook"
	waitingroom "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroom"
	waitingroomevent "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroomevent"
	waitingroomrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroomrules"
	waitingroomsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroomsettings"
	pagesdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/pagesdomain"
	pagesproject "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/pagesproject"
	worker "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/worker"
	workerscrontrigger "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workerscrontrigger"
	workerscustomdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workerscustomdomain"
	workersdeployment "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workersdeployment"
	workersforplatformsdispatchnamespace "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workersforplatformsdispatchnamespace"
	workerskv "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workerskv"
	workerskvnamespace "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workerskvnamespace"
	workersroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workersroute"
	workersscript "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workersscript"
	workersscriptsubdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workersscriptsubdomain"
	workerversion "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workerversion"
	workflow "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workflow"
	zerotrustdevicecustomprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicecustomprofile"
	zerotrustdevicecustomprofilelocaldomainfallback "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicecustomprofilelocaldomainfallback"
	zerotrustdevicedefaultprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicedefaultprofile"
	zerotrustdevicedefaultprofilecertificates "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicedefaultprofilecertificates"
	zerotrustdevicedefaultprofilelocaldomainfallback "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicedefaultprofilelocaldomainfallback"
	zerotrustdevicemanagednetworks "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicemanagednetworks"
	zerotrustdevicepostureintegration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicepostureintegration"
	zerotrustdeviceposturerule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdeviceposturerule"
	zerotrustdevicesettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicesettings"
	zerotrustdextest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdextest"
	zerotrustdlpcustomentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpcustomentry"
	zerotrustdlpcustomprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpcustomprofile"
	zerotrustdlpdataset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpdataset"
	zerotrustdlpentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpentry"
	zerotrustdlpintegrationentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpintegrationentry"
	zerotrustdlppredefinedentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlppredefinedentry"
	zerotrustdlppredefinedprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlppredefinedprofile"
	zerotrustdnslocation "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdnslocation"
	zerotrustgatewaycertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewaycertificate"
	zerotrustgatewaylogging "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewaylogging"
	zerotrustgatewaypolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewaypolicy"
	zerotrustgatewayproxyendpoint "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewayproxyendpoint"
	zerotrustgatewaysettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewaysettings"
	zerotrustlist "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustlist"
	zerotrustnetworkhostnameroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustnetworkhostnameroute"
	zerotrustriskbehavior "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustriskbehavior"
	zerotrustriskscoringintegration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustriskscoringintegration"
	zerotrusttunnelcloudflared "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrusttunnelcloudflared"
	zerotrusttunnelcloudflaredconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrusttunnelcloudflaredconfig"
	zerotrusttunnelcloudflaredroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrusttunnelcloudflaredroute"
	zerotrusttunnelcloudflaredvirtualnetwork "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrusttunnelcloudflaredvirtualnetwork"
	zerotrusttunnelwarpconnector "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrusttunnelwarpconnector"
	zone "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zone"
	zonecachereserve "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonecachereserve"
	zonecachevariants "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonecachevariants"
	zonednssec "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonednssec"
	zonednssettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonednssettings"
	zonehold "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonehold"
	zonelockdown "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonelockdown"
	zonesetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonesetting"
	zonesubscription "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonesubscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessrule.Setup,
		zerotrustaccessaicontrolsmcpportal.Setup,
		zerotrustaccessaicontrolsmcpserver.Setup,
		zerotrustaccessapplication.Setup,
		zerotrustaccesscustompage.Setup,
		zerotrustaccessgroup.Setup,
		zerotrustaccessidentityprovider.Setup,
		zerotrustaccessinfrastructuretarget.Setup,
		zerotrustaccesskeyconfiguration.Setup,
		zerotrustaccessmtlscertificate.Setup,
		zerotrustaccessmtlshostnamesettings.Setup,
		zerotrustaccesspolicy.Setup,
		zerotrustaccessservicetoken.Setup,
		zerotrustaccessshortlivedcertificate.Setup,
		zerotrustaccesstag.Setup,
		zerotrustorganization.Setup,
		account.Setup,
		accountsubscription.Setup,
		accounttoken.Setup,
		apitoken.Setup,
		member.Setup,
		registrardomain.Setup,
		apishield.Setup,
		apishielddiscoveryoperation.Setup,
		apishieldoperation.Setup,
		apishieldoperationschemavalidationsettings.Setup,
		apishieldschema.Setup,
		apishieldschemavalidationsettings.Setup,
		schemavalidationoperationsettings.Setup,
		schemavalidationschemas.Setup,
		schemavalidationsettings.Setup,
		argosmartrouting.Setup,
		argotieredcaching.Setup,
		magicnetworkmonitoringconfiguration.Setup,
		magicnetworkmonitoringrule.Setup,
		magictransitconnector.Setup,
		magictransitsite.Setup,
		magictransitsiteacl.Setup,
		magictransitsitelan.Setup,
		magictransitsitewan.Setup,
		magicwangretunnel.Setup,
		magicwanipsectunnel.Setup,
		magicwanstaticroute.Setup,
		regionaltieredcache.Setup,
		tieredcache.Setup,
		authenticatedoriginpulls.Setup,
		authenticatedoriginpullscertificate.Setup,
		authenticatedoriginpullssettings.Setup,
		certificatepack.Setup,
		customhostname.Setup,
		customhostnamefallbackorigin.Setup,
		customssl.Setup,
		hostnametlssetting.Setup,
		keylesscertificate.Setup,
		mtlscertificate.Setup,
		origincacertificate.Setup,
		totaltls.Setup,
		universalsslsetting.Setup,
		accountdnssettings.Setup,
		accountdnssettingsinternalview.Setup,
		dnsfirewall.Setup,
		dnsrecord.Setup,
		dnszonetransfersacl.Setup,
		dnszonetransfersincoming.Setup,
		dnszonetransfersoutgoing.Setup,
		dnszonetransferspeer.Setup,
		dnszonetransferstsig.Setup,
		emailroutingaddress.Setup,
		emailroutingcatchall.Setup,
		emailroutingdns.Setup,
		emailroutingrule.Setup,
		emailroutingsettings.Setup,
		emailsecurityblocksender.Setup,
		emailsecurityimpersonationregistry.Setup,
		emailsecuritytrusteddomains.Setup,
		botmanagement.Setup,
		custompages.Setup,
		filter.Setup,
		firewallrule.Setup,
		leakedcredentialcheck.Setup,
		leakedcredentialcheckrule.Setup,
		managedtransforms.Setup,
		pagerule.Setup,
		pageshieldpolicy.Setup,
		ratelimit.Setup,
		ruleset.Setup,
		urlnormalizationsettings.Setup,
		useragentblockingrule.Setup,
		healthcheck.Setup,
		list.Setup,
		listitem.Setup,
		loadbalancer.Setup,
		loadbalancermonitor.Setup,
		loadbalancerpool.Setup,
		logpullretention.Setup,
		logpushjob.Setup,
		logpushownershipchallenge.Setup,
		addressmap.Setup,
		byoipprefix.Setup,
		callssfuapp.Setup,
		callsturnapp.Setup,
		cloudconnectorrules.Setup,
		cloudforceonerequest.Setup,
		cloudforceonerequestasset.Setup,
		cloudforceonerequestmessage.Setup,
		cloudforceonerequestpriority.Setup,
		connectivitydirectoryservice.Setup,
		contentscanning.Setup,
		contentscanningexpression.Setup,
		d1database.Setup,
		hyperdriveconfig.Setup,
		image.Setup,
		imagevariant.Setup,
		observatoryscheduledtest.Setup,
		organization.Setup,
		organizationprofile.Setup,
		queue.Setup,
		queueconsumer.Setup,
		regionalhostname.Setup,
		snippet.Setup,
		snippetrules.Setup,
		snippets.Setup,
		ssoconnector.Setup,
		tokenvalidationconfig.Setup,
		tokenvalidationrules.Setup,
		turnstilewidget.Setup,
		user.Setup,
		web3hostname.Setup,
		webanalyticsrule.Setup,
		webanalyticssite.Setup,
		notificationpolicy.Setup,
		notificationpolicywebhooks.Setup,
		providerconfig.Setup,
		r2bucket.Setup,
		r2bucketcors.Setup,
		r2bucketeventnotification.Setup,
		r2bucketlifecycle.Setup,
		r2bucketlock.Setup,
		r2bucketsippy.Setup,
		r2customdomain.Setup,
		r2manageddomain.Setup,
		spectrumapplication.Setup,
		stream.Setup,
		streamaudiotrack.Setup,
		streamcaptionlanguage.Setup,
		streamdownload.Setup,
		streamkey.Setup,
		streamliveinput.Setup,
		streamwatermark.Setup,
		streamwebhook.Setup,
		waitingroom.Setup,
		waitingroomevent.Setup,
		waitingroomrules.Setup,
		waitingroomsettings.Setup,
		pagesdomain.Setup,
		pagesproject.Setup,
		worker.Setup,
		workerscrontrigger.Setup,
		workerscustomdomain.Setup,
		workersdeployment.Setup,
		workersforplatformsdispatchnamespace.Setup,
		workerskv.Setup,
		workerskvnamespace.Setup,
		workersroute.Setup,
		workersscript.Setup,
		workersscriptsubdomain.Setup,
		workerversion.Setup,
		workflow.Setup,
		zerotrustdevicecustomprofile.Setup,
		zerotrustdevicecustomprofilelocaldomainfallback.Setup,
		zerotrustdevicedefaultprofile.Setup,
		zerotrustdevicedefaultprofilecertificates.Setup,
		zerotrustdevicedefaultprofilelocaldomainfallback.Setup,
		zerotrustdevicemanagednetworks.Setup,
		zerotrustdevicepostureintegration.Setup,
		zerotrustdeviceposturerule.Setup,
		zerotrustdevicesettings.Setup,
		zerotrustdextest.Setup,
		zerotrustdlpcustomentry.Setup,
		zerotrustdlpcustomprofile.Setup,
		zerotrustdlpdataset.Setup,
		zerotrustdlpentry.Setup,
		zerotrustdlpintegrationentry.Setup,
		zerotrustdlppredefinedentry.Setup,
		zerotrustdlppredefinedprofile.Setup,
		zerotrustdnslocation.Setup,
		zerotrustgatewaycertificate.Setup,
		zerotrustgatewaylogging.Setup,
		zerotrustgatewaypolicy.Setup,
		zerotrustgatewayproxyendpoint.Setup,
		zerotrustgatewaysettings.Setup,
		zerotrustlist.Setup,
		zerotrustnetworkhostnameroute.Setup,
		zerotrustriskbehavior.Setup,
		zerotrustriskscoringintegration.Setup,
		zerotrusttunnelcloudflared.Setup,
		zerotrusttunnelcloudflaredconfig.Setup,
		zerotrusttunnelcloudflaredroute.Setup,
		zerotrusttunnelcloudflaredvirtualnetwork.Setup,
		zerotrusttunnelwarpconnector.Setup,
		zone.Setup,
		zonecachereserve.Setup,
		zonecachevariants.Setup,
		zonednssec.Setup,
		zonednssettings.Setup,
		zonehold.Setup,
		zonelockdown.Setup,
		zonesetting.Setup,
		zonesubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessrule.SetupGated,
		zerotrustaccessaicontrolsmcpportal.SetupGated,
		zerotrustaccessaicontrolsmcpserver.SetupGated,
		zerotrustaccessapplication.SetupGated,
		zerotrustaccesscustompage.SetupGated,
		zerotrustaccessgroup.SetupGated,
		zerotrustaccessidentityprovider.SetupGated,
		zerotrustaccessinfrastructuretarget.SetupGated,
		zerotrustaccesskeyconfiguration.SetupGated,
		zerotrustaccessmtlscertificate.SetupGated,
		zerotrustaccessmtlshostnamesettings.SetupGated,
		zerotrustaccesspolicy.SetupGated,
		zerotrustaccessservicetoken.SetupGated,
		zerotrustaccessshortlivedcertificate.SetupGated,
		zerotrustaccesstag.SetupGated,
		zerotrustorganization.SetupGated,
		account.SetupGated,
		accountsubscription.SetupGated,
		accounttoken.SetupGated,
		apitoken.SetupGated,
		member.SetupGated,
		registrardomain.SetupGated,
		apishield.SetupGated,
		apishielddiscoveryoperation.SetupGated,
		apishieldoperation.SetupGated,
		apishieldoperationschemavalidationsettings.SetupGated,
		apishieldschema.SetupGated,
		apishieldschemavalidationsettings.SetupGated,
		schemavalidationoperationsettings.SetupGated,
		schemavalidationschemas.SetupGated,
		schemavalidationsettings.SetupGated,
		argosmartrouting.SetupGated,
		argotieredcaching.SetupGated,
		magicnetworkmonitoringconfiguration.SetupGated,
		magicnetworkmonitoringrule.SetupGated,
		magictransitconnector.SetupGated,
		magictransitsite.SetupGated,
		magictransitsiteacl.SetupGated,
		magictransitsitelan.SetupGated,
		magictransitsitewan.SetupGated,
		magicwangretunnel.SetupGated,
		magicwanipsectunnel.SetupGated,
		magicwanstaticroute.SetupGated,
		regionaltieredcache.SetupGated,
		tieredcache.SetupGated,
		authenticatedoriginpulls.SetupGated,
		authenticatedoriginpullscertificate.SetupGated,
		authenticatedoriginpullssettings.SetupGated,
		certificatepack.SetupGated,
		customhostname.SetupGated,
		customhostnamefallbackorigin.SetupGated,
		customssl.SetupGated,
		hostnametlssetting.SetupGated,
		keylesscertificate.SetupGated,
		mtlscertificate.SetupGated,
		origincacertificate.SetupGated,
		totaltls.SetupGated,
		universalsslsetting.SetupGated,
		accountdnssettings.SetupGated,
		accountdnssettingsinternalview.SetupGated,
		dnsfirewall.SetupGated,
		dnsrecord.SetupGated,
		dnszonetransfersacl.SetupGated,
		dnszonetransfersincoming.SetupGated,
		dnszonetransfersoutgoing.SetupGated,
		dnszonetransferspeer.SetupGated,
		dnszonetransferstsig.SetupGated,
		emailroutingaddress.SetupGated,
		emailroutingcatchall.SetupGated,
		emailroutingdns.SetupGated,
		emailroutingrule.SetupGated,
		emailroutingsettings.SetupGated,
		emailsecurityblocksender.SetupGated,
		emailsecurityimpersonationregistry.SetupGated,
		emailsecuritytrusteddomains.SetupGated,
		botmanagement.SetupGated,
		custompages.SetupGated,
		filter.SetupGated,
		firewallrule.SetupGated,
		leakedcredentialcheck.SetupGated,
		leakedcredentialcheckrule.SetupGated,
		managedtransforms.SetupGated,
		pagerule.SetupGated,
		pageshieldpolicy.SetupGated,
		ratelimit.SetupGated,
		ruleset.SetupGated,
		urlnormalizationsettings.SetupGated,
		useragentblockingrule.SetupGated,
		healthcheck.SetupGated,
		list.SetupGated,
		listitem.SetupGated,
		loadbalancer.SetupGated,
		loadbalancermonitor.SetupGated,
		loadbalancerpool.SetupGated,
		logpullretention.SetupGated,
		logpushjob.SetupGated,
		logpushownershipchallenge.SetupGated,
		addressmap.SetupGated,
		byoipprefix.SetupGated,
		callssfuapp.SetupGated,
		callsturnapp.SetupGated,
		cloudconnectorrules.SetupGated,
		cloudforceonerequest.SetupGated,
		cloudforceonerequestasset.SetupGated,
		cloudforceonerequestmessage.SetupGated,
		cloudforceonerequestpriority.SetupGated,
		connectivitydirectoryservice.SetupGated,
		contentscanning.SetupGated,
		contentscanningexpression.SetupGated,
		d1database.SetupGated,
		hyperdriveconfig.SetupGated,
		image.SetupGated,
		imagevariant.SetupGated,
		observatoryscheduledtest.SetupGated,
		organization.SetupGated,
		organizationprofile.SetupGated,
		queue.SetupGated,
		queueconsumer.SetupGated,
		regionalhostname.SetupGated,
		snippet.SetupGated,
		snippetrules.SetupGated,
		snippets.SetupGated,
		ssoconnector.SetupGated,
		tokenvalidationconfig.SetupGated,
		tokenvalidationrules.SetupGated,
		turnstilewidget.SetupGated,
		user.SetupGated,
		web3hostname.SetupGated,
		webanalyticsrule.SetupGated,
		webanalyticssite.SetupGated,
		notificationpolicy.SetupGated,
		notificationpolicywebhooks.SetupGated,
		providerconfig.SetupGated,
		r2bucket.SetupGated,
		r2bucketcors.SetupGated,
		r2bucketeventnotification.SetupGated,
		r2bucketlifecycle.SetupGated,
		r2bucketlock.SetupGated,
		r2bucketsippy.SetupGated,
		r2customdomain.SetupGated,
		r2manageddomain.SetupGated,
		spectrumapplication.SetupGated,
		stream.SetupGated,
		streamaudiotrack.SetupGated,
		streamcaptionlanguage.SetupGated,
		streamdownload.SetupGated,
		streamkey.SetupGated,
		streamliveinput.SetupGated,
		streamwatermark.SetupGated,
		streamwebhook.SetupGated,
		waitingroom.SetupGated,
		waitingroomevent.SetupGated,
		waitingroomrules.SetupGated,
		waitingroomsettings.SetupGated,
		pagesdomain.SetupGated,
		pagesproject.SetupGated,
		worker.SetupGated,
		workerscrontrigger.SetupGated,
		workerscustomdomain.SetupGated,
		workersdeployment.SetupGated,
		workersforplatformsdispatchnamespace.SetupGated,
		workerskv.SetupGated,
		workerskvnamespace.SetupGated,
		workersroute.SetupGated,
		workersscript.SetupGated,
		workersscriptsubdomain.SetupGated,
		workerversion.SetupGated,
		workflow.SetupGated,
		zerotrustdevicecustomprofile.SetupGated,
		zerotrustdevicecustomprofilelocaldomainfallback.SetupGated,
		zerotrustdevicedefaultprofile.SetupGated,
		zerotrustdevicedefaultprofilecertificates.SetupGated,
		zerotrustdevicedefaultprofilelocaldomainfallback.SetupGated,
		zerotrustdevicemanagednetworks.SetupGated,
		zerotrustdevicepostureintegration.SetupGated,
		zerotrustdeviceposturerule.SetupGated,
		zerotrustdevicesettings.SetupGated,
		zerotrustdextest.SetupGated,
		zerotrustdlpcustomentry.SetupGated,
		zerotrustdlpcustomprofile.SetupGated,
		zerotrustdlpdataset.SetupGated,
		zerotrustdlpentry.SetupGated,
		zerotrustdlpintegrationentry.SetupGated,
		zerotrustdlppredefinedentry.SetupGated,
		zerotrustdlppredefinedprofile.SetupGated,
		zerotrustdnslocation.SetupGated,
		zerotrustgatewaycertificate.SetupGated,
		zerotrustgatewaylogging.SetupGated,
		zerotrustgatewaypolicy.SetupGated,
		zerotrustgatewayproxyendpoint.SetupGated,
		zerotrustgatewaysettings.SetupGated,
		zerotrustlist.SetupGated,
		zerotrustnetworkhostnameroute.SetupGated,
		zerotrustriskbehavior.SetupGated,
		zerotrustriskscoringintegration.SetupGated,
		zerotrusttunnelcloudflared.SetupGated,
		zerotrusttunnelcloudflaredconfig.SetupGated,
		zerotrusttunnelcloudflaredroute.SetupGated,
		zerotrusttunnelcloudflaredvirtualnetwork.SetupGated,
		zerotrusttunnelwarpconnector.SetupGated,
		zone.SetupGated,
		zonecachereserve.SetupGated,
		zonecachevariants.SetupGated,
		zonednssec.SetupGated,
		zonednssettings.SetupGated,
		zonehold.SetupGated,
		zonelockdown.SetupGated,
		zonesetting.SetupGated,
		zonesubscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
