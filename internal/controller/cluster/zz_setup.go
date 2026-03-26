// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accessrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/accessrule"
	zerotrustaccessaicontrolsmcpportal "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessaicontrolsmcpportal"
	zerotrustaccessaicontrolsmcpserver "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessaicontrolsmcpserver"
	zerotrustaccessapplication "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessapplication"
	zerotrustaccesscustompage "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccesscustompage"
	zerotrustaccessgroup "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessgroup"
	zerotrustaccessidentityprovider "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessidentityprovider"
	zerotrustaccessinfrastructuretarget "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessinfrastructuretarget"
	zerotrustaccesskeyconfiguration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccesskeyconfiguration"
	zerotrustaccessmtlscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessmtlscertificate"
	zerotrustaccessmtlshostnamesettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessmtlshostnamesettings"
	zerotrustaccesspolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccesspolicy"
	zerotrustaccessservicetoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessservicetoken"
	zerotrustaccessshortlivedcertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccessshortlivedcertificate"
	zerotrustaccesstag "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustaccesstag"
	zerotrustorganization "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/access/zerotrustorganization"
	account "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/account/account"
	accountsubscription "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/account/accountsubscription"
	accounttoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/account/accounttoken"
	apitoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/account/apitoken"
	member "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/account/member"
	registrardomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/account/registrardomain"
	apishield "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/apishield"
	apishielddiscoveryoperation "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/apishielddiscoveryoperation"
	apishieldoperation "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/apishieldoperation"
	apishieldoperationschemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/apishieldoperationschemavalidationsettings"
	apishieldschema "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/apishieldschema"
	apishieldschemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/apishieldschemavalidationsettings"
	schemavalidationoperationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/schemavalidationoperationsettings"
	schemavalidationschemas "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/schemavalidationschemas"
	schemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/apishield/schemavalidationsettings"
	argosmartrouting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/argosmartrouting"
	argotieredcaching "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/argotieredcaching"
	magicnetworkmonitoringconfiguration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magicnetworkmonitoringconfiguration"
	magicnetworkmonitoringrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magicnetworkmonitoringrule"
	magictransitconnector "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magictransitconnector"
	magictransitsite "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magictransitsite"
	magictransitsiteacl "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magictransitsiteacl"
	magictransitsitelan "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magictransitsitelan"
	magictransitsitewan "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magictransitsitewan"
	magicwangretunnel "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magicwangretunnel"
	magicwanipsectunnel "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magicwanipsectunnel"
	magicwanstaticroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/magicwanstaticroute"
	regionaltieredcache "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/regionaltieredcache"
	tieredcache "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/argo/tieredcache"
	authenticatedoriginpulls "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/authenticatedoriginpulls"
	authenticatedoriginpullscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/authenticatedoriginpullscertificate"
	authenticatedoriginpullssettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/authenticatedoriginpullssettings"
	certificatepack "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/certificatepack"
	customhostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/customhostname"
	customhostnamefallbackorigin "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/customhostnamefallbackorigin"
	customssl "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/customssl"
	hostnametlssetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/hostnametlssetting"
	keylesscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/keylesscertificate"
	mtlscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/mtlscertificate"
	origincacertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/origincacertificate"
	totaltls "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/totaltls"
	universalsslsetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/certificate/universalsslsetting"
	accountdnssettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/accountdnssettings"
	accountdnssettingsinternalview "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/accountdnssettingsinternalview"
	dnsfirewall "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/dnsfirewall"
	dnsrecord "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/dnsrecord"
	dnszonetransfersacl "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/dnszonetransfersacl"
	dnszonetransfersincoming "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/dnszonetransfersincoming"
	dnszonetransfersoutgoing "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/dnszonetransfersoutgoing"
	dnszonetransferspeer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/dnszonetransferspeer"
	dnszonetransferstsig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/dns/dnszonetransferstsig"
	emailroutingaddress "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailroutingaddress"
	emailroutingcatchall "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailroutingcatchall"
	emailroutingdns "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailroutingdns"
	emailroutingrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailroutingrule"
	emailroutingsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailroutingsettings"
	emailsecurityblocksender "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailsecurityblocksender"
	emailsecurityimpersonationregistry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailsecurityimpersonationregistry"
	emailsecuritytrusteddomains "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/email/emailsecuritytrusteddomains"
	botmanagement "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/botmanagement"
	custompages "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/custompages"
	filter "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/filter"
	firewallrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/firewallrule"
	leakedcredentialcheck "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/leakedcredentialcheck"
	leakedcredentialcheckrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/leakedcredentialcheckrule"
	managedtransforms "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/managedtransforms"
	pagerule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/pagerule"
	pageshieldpolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/pageshieldpolicy"
	ratelimit "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/ratelimit"
	ruleset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/ruleset"
	urlnormalizationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/urlnormalizationsettings"
	useragentblockingrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/firewall/useragentblockingrule"
	healthcheck "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/healthcheck/healthcheck"
	list "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/list/list"
	listitem "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/list/listitem"
	loadbalancer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/loadbalancer/loadbalancer"
	loadbalancermonitor "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/loadbalancer/loadbalancermonitor"
	loadbalancerpool "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/loadbalancer/loadbalancerpool"
	logpullretention "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/logpush/logpullretention"
	logpushjob "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/logpush/logpushjob"
	logpushownershipchallenge "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/logpush/logpushownershipchallenge"
	addressmap "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/addressmap"
	byoipprefix "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/byoipprefix"
	callssfuapp "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/callssfuapp"
	callsturnapp "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/callsturnapp"
	cloudconnectorrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/cloudconnectorrules"
	cloudforceonerequest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/cloudforceonerequest"
	cloudforceonerequestasset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/cloudforceonerequestasset"
	cloudforceonerequestmessage "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/cloudforceonerequestmessage"
	cloudforceonerequestpriority "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/cloudforceonerequestpriority"
	connectivitydirectoryservice "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/connectivitydirectoryservice"
	contentscanning "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/contentscanning"
	contentscanningexpression "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/contentscanningexpression"
	d1database "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/d1database"
	hyperdriveconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/hyperdriveconfig"
	image "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/image"
	imagevariant "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/imagevariant"
	observatoryscheduledtest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/observatoryscheduledtest"
	organization "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/organization"
	organizationprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/organizationprofile"
	queue "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/queue"
	queueconsumer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/queueconsumer"
	regionalhostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/regionalhostname"
	snippet "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/snippet"
	snippetrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/snippetrules"
	snippets "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/snippets"
	ssoconnector "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/ssoconnector"
	tokenvalidationconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/tokenvalidationconfig"
	tokenvalidationrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/tokenvalidationrules"
	turnstilewidget "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/turnstilewidget"
	user "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/user"
	web3hostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/web3hostname"
	webanalyticsrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/webanalyticsrule"
	webanalyticssite "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/misc/webanalyticssite"
	notificationpolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/notification/notificationpolicy"
	notificationpolicywebhooks "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/notification/notificationpolicywebhooks"
	providerconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/providerconfig"
	r2bucket "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2bucket"
	r2bucketcors "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2bucketcors"
	r2bucketeventnotification "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2bucketeventnotification"
	r2bucketlifecycle "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2bucketlifecycle"
	r2bucketlock "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2bucketlock"
	r2bucketsippy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2bucketsippy"
	r2customdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2customdomain"
	r2manageddomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/r2/r2manageddomain"
	spectrumapplication "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/spectrum/spectrumapplication"
	stream "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/stream"
	streamaudiotrack "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/streamaudiotrack"
	streamcaptionlanguage "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/streamcaptionlanguage"
	streamdownload "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/streamdownload"
	streamkey "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/streamkey"
	streamliveinput "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/streamliveinput"
	streamwatermark "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/streamwatermark"
	streamwebhook "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/stream/streamwebhook"
	waitingroom "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/waitingroom/waitingroom"
	waitingroomevent "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/waitingroom/waitingroomevent"
	waitingroomrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/waitingroom/waitingroomrules"
	waitingroomsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/waitingroom/waitingroomsettings"
	pagesdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/pagesdomain"
	pagesproject "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/pagesproject"
	worker "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/worker"
	workerscrontrigger "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workerscrontrigger"
	workerscustomdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workerscustomdomain"
	workersdeployment "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workersdeployment"
	workersforplatformsdispatchnamespace "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workersforplatformsdispatchnamespace"
	workerskv "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workerskv"
	workerskvnamespace "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workerskvnamespace"
	workersroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workersroute"
	workersscript "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workersscript"
	workersscriptsubdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workersscriptsubdomain"
	workerversion "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workerversion"
	workflow "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/workers/workflow"
	zerotrustdevicecustomprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicecustomprofile"
	zerotrustdevicecustomprofilelocaldomainfallback "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicecustomprofilelocaldomainfallback"
	zerotrustdevicedefaultprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicedefaultprofile"
	zerotrustdevicedefaultprofilecertificates "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicedefaultprofilecertificates"
	zerotrustdevicedefaultprofilelocaldomainfallback "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicedefaultprofilelocaldomainfallback"
	zerotrustdevicemanagednetworks "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicemanagednetworks"
	zerotrustdevicepostureintegration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicepostureintegration"
	zerotrustdeviceposturerule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdeviceposturerule"
	zerotrustdevicesettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdevicesettings"
	zerotrustdextest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdextest"
	zerotrustdlpcustomentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdlpcustomentry"
	zerotrustdlpcustomprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdlpcustomprofile"
	zerotrustdlpdataset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdlpdataset"
	zerotrustdlpentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdlpentry"
	zerotrustdlpintegrationentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdlpintegrationentry"
	zerotrustdlppredefinedentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdlppredefinedentry"
	zerotrustdlppredefinedprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdlppredefinedprofile"
	zerotrustdnslocation "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustdnslocation"
	zerotrustgatewaycertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustgatewaycertificate"
	zerotrustgatewaylogging "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustgatewaylogging"
	zerotrustgatewaypolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustgatewaypolicy"
	zerotrustgatewayproxyendpoint "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustgatewayproxyendpoint"
	zerotrustgatewaysettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustgatewaysettings"
	zerotrustlist "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustlist"
	zerotrustnetworkhostnameroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustnetworkhostnameroute"
	zerotrustriskbehavior "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustriskbehavior"
	zerotrustriskscoringintegration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrustriskscoringintegration"
	zerotrusttunnelcloudflared "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrusttunnelcloudflared"
	zerotrusttunnelcloudflaredconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrusttunnelcloudflaredconfig"
	zerotrusttunnelcloudflaredroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrusttunnelcloudflaredroute"
	zerotrusttunnelcloudflaredvirtualnetwork "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrusttunnelcloudflaredvirtualnetwork"
	zerotrusttunnelwarpconnector "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zerotrust/zerotrusttunnelwarpconnector"
	zone "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zone"
	zonecachereserve "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonecachereserve"
	zonecachevariants "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonecachevariants"
	zonednssec "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonednssec"
	zonednssettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonednssettings"
	zonehold "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonehold"
	zonelockdown "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonelockdown"
	zonesetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonesetting"
	zonesubscription "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/cluster/zone/zonesubscription"
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
